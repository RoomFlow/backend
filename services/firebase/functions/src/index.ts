import * as functions from 'firebase-functions';
import * as admin from 'firebase-admin';

admin.initializeApp();
const db = admin.firestore();

interface Update {
    timestamp?: Date;
    motion?: Boolean;
    light?: String;
    sound?: String;
}

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
    const { timestamp, motion, sound, light } = content;
    // const { deviceId } = attributes;

    const updateObj: Update = {};

    if (timestamp) updateObj.timestamp = new Date(timestamp * 1000);
    if (motion) updateObj.motion = motion;
    if (light) updateObj.light = light;
    if (sound) updateObj.sound = sound;

    db.collection("rooms").doc(room).update(updateObj)
    .then(() => {
        console.log("Document successfully updated!");
    })
    .catch((error) => {
        console.error("Error updating document: ", error);
    });
});