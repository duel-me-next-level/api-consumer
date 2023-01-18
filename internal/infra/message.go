package message

// SendMQTTMessage sends a message over the MQTT protocol
func SendMQTTMessage(topic string, message []byte) error {
	// Connect to the MQTT broker
	client, err := mqtt.NewClient(mqtt.ClientOptions{
		// Options go here
	})
	if err != nil {
		return err
	}
	defer client.Disconnect()

	// Publish the message to the specified topic
	token := client.Publish(topic, 0, false, message)
	token.Wait()
	if token.Error() != nil {
		return token.Error()
	}

	return nil
}
