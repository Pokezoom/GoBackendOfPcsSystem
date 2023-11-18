from scipy.spatial import distance as dist
from collections import OrderedDict
import numpy as np
import time
import dlib
import cv2

FACIAL_LANDMARKS_IDXS = OrderedDict([
    ("mouth",(48,68)),
     ("right_eyebrow",(17,22)),
     ("left_eyebrow",(22,27)),
     ("right_eye",(36,42)),
     ("left_eye",(42,48)),
     ("nose",(27,36)),
     ("jaw",(0,17))
])

def eye_aspect_ratio(eye):
    A = dist.euclidean(eye[1], eye[5])
    B = dist.euclidean(eye[2], eye[4])
    C = dist.euclidean(eye[0], eye[3])
    ear = (A+B)/(2.0*C)
    return ear

# Add MAR calculation function
def mouth_aspect_ratio(mouth):
    A = dist.euclidean(mouth[13], mouth[19]) # 49-55
    B = dist.euclidean(mouth[14], mouth[18]) # 50-56
    C = dist.euclidean(mouth[15], mouth[17]) # 51-57
    D = dist.euclidean(mouth[12], mouth[16]) # 48-54
    mar = (A + B + C) / (2.0 * D)
    return mar

# Define parameters for yawning
MAR_THRESH = 0.6 # This is an example value, adjust based on testing
YAWN_CONSEC_FRAMES = 3

# Variables for yawning
yawn_counter = 0
yawns = 0


EYE_AR_THRESH = 0.3
EYE_AR_CONSEC_FRAME = 3


COUNTER = 0
TOTAL = 0

print("[INFO] loading facial landmark predictor...")
detector = dlib.get_frontal_face_detector()
predictor = dlib.shape_predictor('shape_predictor_68_face_landmarks.dat')

(lStart, lEnd) = FACIAL_LANDMARKS_IDXS["left_eye"]
(rStart, rEnd) = FACIAL_LANDMARKS_IDXS["right_eye"]

print("[INFO] starting video stream thread...")
vs = cv2.VideoCapture('test1.mp4')
time.sleep(1.0)

def shape_to_np(shape, dtype="int"):
    coords = np.zeros((shape.num_parts, 2), dtype=dtype)
    for i in range(0, shape.num_parts):
        coords[i] = (shape.part(i).x, shape.part(i).y)
    return coords

while True:
    frame = vs.read()[1]
    if frame is None:
        break
    (h, w) = frame.shape[:2]
    width = 1200
    r = width / float(w)
    dim = (width, int(h * r))
    frame = cv2.resize(frame, dim, interpolation=cv2.INTER_AREA)
    gray = cv2.cvtColor(frame, cv2.COLOR_BGR2GRAY)

    rects = detector(gray, 0)

    for rect in rects:
        shape = predictor(gray, rect)
        shape = shape_to_np(shape)

        leftEye = shape[lStart:lEnd]
        rightEye = shape[rStart:rEnd]
        leftEAR = eye_aspect_ratio(leftEye)
        rightEAR = eye_aspect_ratio(rightEye)
        # Process mouth for yawning
        mouth = shape[48:68]  # Mouth landmarks
        mar = mouth_aspect_ratio(mouth)
        ear = (leftEAR + rightEAR) / 2.0

        mouthHull = cv2.convexHull(mouth)
        cv2.drawContours(frame, [mouthHull], -1, (0, 255, 0), 1)

        leftEyeHull = cv2.convexHull(leftEye)
        rightEyeHull = cv2.convexHull(rightEye)
        cv2.drawContours(frame, [leftEyeHull], -1, (0, 255, 0), 1)
        cv2.drawContours(frame, [rightEyeHull], -1, (0, 255, 0), 1)

        if ear < EYE_AR_THRESH:
            COUNTER += 1
        else:
            if COUNTER >= EYE_AR_CONSEC_FRAME:
                TOTAL += 1

            COUNTER = 0

        if mar > MAR_THRESH:
            yawn_counter += 1
        else:
            if yawn_counter >= YAWN_CONSEC_FRAMES:
                yawns += 1
            yawn_counter = 0

        cv2.putText(frame, "Blinks: {}".format(TOTAL), (10, 30),
                    cv2.FONT_HERSHEY_SIMPLEX, 0.7, (0, 0, 255), 2)
        cv2.putText(frame, "EAR: {:.2f}".format(ear), (300, 30),
                    cv2.FONT_HERSHEY_SIMPLEX, 0.7, (0, 0, 255), 2)
        cv2.putText(frame, "Yawns: {}".format(yawns), (10, 60),
                    cv2.FONT_HERSHEY_SIMPLEX, 0.7, (0, 0, 255), 2)

    cv2.imshow("Frame", frame)
    key = cv2.waitKey(10) & 0xFF

    if key == 27:
        break

vs.release()
cv2.destroyAllWindows()
