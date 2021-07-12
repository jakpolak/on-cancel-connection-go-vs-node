const express = require("express");

async function main() {
    const app = express();

    app.get("/cancel", handlerWithRequestCancellation);
    app.listen(3000, "127.0.0.1", () => {
        console.log("Listening port 3000");
    });
}

async function handlerWithRequestCancellation(req, res) {
    let requestCancelled = false;

    req.connection.on("close", () => {
        console.log("Request has been cancelled");
        requestCancelled = true;
    });

    const words = ["Hello", "Node", "JS", "Connection", "Canceling"];

    for (const word of words) {
        await slowConsoleLog(word, requestCancelled);
        if (requestCancelled) {
            return;
        }
    }

    res.send("OK");
}

async function slowConsoleLog(word) {
    console.log(`Processing: ${word}`);
    await sleep(3000);
}

function sleep(millis) {
    return new Promise((resolve) => setTimeout(resolve, millis));
}

main();
