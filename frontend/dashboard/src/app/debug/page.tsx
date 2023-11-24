"use client";
import { createPromiseClient } from "@connectrpc/connect";
import { StateManagerService } from "@/proto/state/v1/state_connectweb";
import { createConnectTransport } from "@bufbuild/connect-web";
import { GetBlockStatesRequest } from "@/proto/state/v1/block_pb";
import { useQuery } from '@tanstack/react-query';

export default function Test() {
    const transport = createConnectTransport({
        // baseUrl: process.env.NEXT_PUBLIC_API_ENDPOINT!,
        baseUrl: "http://localhost:8080",
    });
    const sendData = () => {
        (async () => {
            const client = createPromiseClient(StateManagerService, transport);
            const res = await client.getBlockStates(
                new GetBlockStatesRequest({}),
            );
            console.log(res);
        })();
    };

    const {trains} =

    return (
        <div>
            <h1>デバッグ画面</h1>

            <button onClick={sendData}>Test</button>
            {}
        </div>
    );
}
