"use client";
import {createPromiseClient} from "@connectrpc/connect";
import {StateManagerService} from "@/proto/state/v1/state_connectweb";
import {createConnectTransport} from "@bufbuild/connect-web";

export default function Test() {

    const transport = createConnectTransport(
        {
            // baseUrl: process.env.NEXT_PUBLIC_API_ENDPOINT!,
            baseUrl: "http://localhost:8080",
        });
    const sendData = () => {
        (async () => {
            const client = createPromiseClient(StateManagerService, transport);
            const res = await client.getBlockStates({});
            console.log(res);
        })();
    }

    return (
        <div>
            <h1>Test</h1>
            <button onClick={sendData}>Test</button>
        </div>
    )
}