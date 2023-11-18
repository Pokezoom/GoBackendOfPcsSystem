import dlib
import cv2
import numpy as np
from scipy.spatial import distance as dist
from collections import OrderedDict

# Initialize the face detector and facial landmark predictor
detector = dlib.get_frontal_face_detector()
predictor = dlib.shape_predictor('shape_predictor_68_face_landmarks.dat')

# Define the facial landmarks for the eyes and mouth
FACIAL_LANDMARKS_IDXS = OrderedDict([
    ("mouth", (48, 68)),
    ("right_eyebrow", (17, 22)),
    ("left_eyebrow", (22, 27)),
    ("right_eye", (36, 42)),
    ("left_eye", (42, 48)),
    ("nose", (27, 36)),
    ("jaw", (0, 17))
])

# Function to convert shape to numpy array
def shape_to_np(shape, dtype="int"):
    coords = np.zeros((shape.num_parts, 2), dtype=dtype)
    for i in range(0, shape.num_parts):
        coords[i] = (shape.part(i).x, shape.part(i).y)
    return coords

# Function to calculate Eye Aspect Ratio (EAR)
def eye_aspect_ratio(eye):
    A = dist.euclidean(eye[1], eye[5])
    B = dist.euclidean(eye[2], eye[4])
    C = dist.euclidean(eye[0], eye[3])
    ear = (A + B) / (2.0 * C)
    return ear

# Function to calculate Mouth Aspect Ratio (MAR)
def mouth_aspect_ratio(mouth):
    A = dist.euclidean(mouth[13], mouth[19])
    B = dist.euclidean(mouth[14], mouth[18])
    C = dist.euclidean(mouth[15], mouth[17])
    D = dist.euclidean(mouth[12], mouth[16])
    mar = (A + B + C) / (2.0 * D)
    return mar

# Thresholds for detecting blinks and yawns
EYE_AR_THRESH = 0.3
MAR_THRESH = 0.6

# Load an image
image_path = 'image.jpg'  # Replace with your image path
image = cv2.imread(image_path)
if image is None:
    raise ValueError(f"Image not found at the path: {image_path}")
gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)

# Upsample the image before running the detector
# Increase the second argument to detect smaller faces; for example, 2 or higher
gray = cv2.cvtColor(image, cv2.COLOR_BGR2GRAY)

# Detect faces in the grayscale image
rects = detector(gray, 2)

# Process each face detected
for rect in rects:
    # Scale the face coordinates back down to the original image size
    scaled_rect = dlib.rectangle(
        int(rect.left() / 2), int(rect.top() / 2),
        int(rect.right() / 2), int(rect.bottom() / 2)
    )

    shape = predictor(gray, rect)
    shape_np = shape_to_np(shape)

    # Extract the left and right eye coordinates
    leftEye = shape_np[FACIAL_LANDMARKS_IDXS["left_eye"][0]:FACIAL_LANDMARKS_IDXS["left_eye"][1]]
    rightEye = shape_np[FACIAL_LANDMARKS_IDXS["right_eye"][0]:FACIAL_LANDMARKS_IDXS["right_eye"][1]]
    mouth = shape_np[FACIAL_LANDMARKS_IDXS["mouth"][0]:FACIAL_LANDMARKS_IDXS["mouth"][1]]

    # Calculate EAR and MAR
    ear = (eye_aspect_ratio(leftEye) + eye_aspect_ratio(rightEye)) / 2.0
    mar = mouth_aspect_ratio(mouth)

    # Draw contours around the eyes and mouth
    cv2.drawContours(image, [cv2.convexHull(leftEye)], -1, (0, 255, 0), 1)
    cv2.drawContours(image, [cv2.convexHull(rightEye)], -1, (0, 255, 0), 1)
    cv2.drawContours(image, [cv2.convexHull(mouth)], -1, (0, 255, 0), 1)

    # Display EAR and MAR on the image near the detected face
    ear_text = f"EAR: {ear:.2f}"
    mar_text = f"MAR: {mar:.2f}"
    cv2.putText(image, ear_text, (rect.left(), rect.top() - 20), cv2.FONT_HERSHEY_SIMPLEX, 0.5, (0, 0, 255), 2)
    cv2.putText(image, mar_text, (rect.left(), rect.top() - 40), cv2.FONT_HERSHEY_SIMPLEX, 0.5, (0, 0, 255), 2)

# Display the image with detections
cv2.imshow("Output", image)
cv2.waitKey(0)
cv2.destroyAllWindows()