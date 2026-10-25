print("AI started\n")

# Model parameter
weight = 0.37

# One experience
input = 1
target = 3

# Prediction
prediction = input * weight

print("Input:", input)
print("Target:", target)
print("Prediction:", prediction)

# Error
error = target - prediction

# Loss
loss = error ** 2

print("Error:", error)
print("Loss:", loss)

# Gradient
gradient = -2 * input * error

print("Gradient:", gradient)

# Learning
learning_rate = 0.001
weight = weight - learning_rate * gradient

print("\nUpdated weight:", weight)