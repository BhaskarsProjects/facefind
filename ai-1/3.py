print("AI started\n")

# Model parameter
weight = 0.37

# Controls how quickly the model learns
learning_rate = 0.001

# Experiences
inputs = [1, 2, 3, 4, 5]
targets = [3, 6, 9, 12, 15]

data_size = 5

# Show initial weight
print("Initial weight:", weight)

# Training
for epoch in range(1000):

    total_loss = 0.0

    for i in range(data_size):

        input = inputs[i]
        target = targets[i]

        # Forward pass
        prediction = input * weight

        # Show first prediction
        if epoch == 0 and i == 0:
            print("First input:", input)
            print("First target:", target)
            print("First prediction:", prediction)

        # Error
        error = target - prediction

        # Loss
        loss = error ** 2
        total_loss += loss

        # Gradient
        gradient = -2 * input * error

        # Gradient descent
        weight = weight - learning_rate * gradient

    if epoch % 100 == 0:
        print(
            f"Epoch: {epoch}"
            f" | Loss: {total_loss}"
            f" | Weight: {weight}"
        )

print("\nTraining finished")
print("Learned weight:", weight)

# Give the model something it has NEVER seen
new_input = 10

prediction = new_input * weight

print("\nNew input:", new_input)
print("Prediction:", prediction)