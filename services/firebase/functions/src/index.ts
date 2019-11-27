import * as functions from 'firebase-functions';
import * as admin from 'firebase-admin';

admin.initializeApp();
const db = admin.firestore();

/**
 * Background Cloud Function to be triggered by Pub/Sub.
 * This function gets executed when telemetry data gets
 * send to IoT Core and consequently a Pub/Sub message
 * gets published to the selected topic.
 *
 */
exports.telemetryToFirestore = functions.pubsub.topic('atest-pub').onPublish((message) => {
    const telemetry = message.json;
    // const attributes = message.attributes;

    if (!telemetry.room || !telemetry.content) {
        throw new Error('No telemetry data was provided!');
    }

    const { room, content } = telemetry;
    // const { deviceId } = attributes;

    // Add telemetry data to new document in collection of same name as room
    db.collection(room).add({
        timestamp: new Date(content.timestamp * 1000),
        motion: content.motion,
    }).then((writeResult) => {
        console.log({ 'result': 'Message with ID: ' + writeResult.id + ' added.' });
        return;
    }).catch((err) => {
        console.log(err);
        return;
    });
});