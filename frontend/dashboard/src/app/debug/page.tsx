"use client";
import {createPromiseClient} from "@connectrpc/connect";
import {createConnectTransport} from "@bufbuild/connect-web";
import {GetBlockStatesRequest} from "@/proto/state/v1/block_pb";
import {useQuery} from '@tanstack/react-query';
import {StateManagerService} from "@/proto/state/v1/state-StateManagerService_connectquery"
import {Client} from "@/components/Client";
import {TrainList} from "@/components/debug/TrainList";

export default function Test() {
    const baseUrl = "http://localhost:8080";

    return (
        <Client baseUrl={baseUrl}>
            <div>
                <h1>デバッグ画面</h1>
                <TrainList />
            </div>
        </Client>
    );
}
