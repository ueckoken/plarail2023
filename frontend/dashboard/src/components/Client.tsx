import {TransportProvider} from "@connectrpc/connect-query";
import {createConnectTransport} from "@connectrpc/connect-web";
import {QueryClient, QueryClientProvider} from "@tanstack/react-query";
import {FC, ReactNode} from "react";

type Props = {
    baseUrl: string;
    children: ReactNode;
};

export const Client: FC<Props> = ({baseUrl, children}) => {
    const transport = createConnectTransport({
        baseUrl: baseUrl,
    });
    const client = new QueryClient();

    return (
        <TransportProvider transport={transport}>
            <QueryClientProvider client={client}>{children}</QueryClientProvider>
        </TransportProvider>
    );
};
