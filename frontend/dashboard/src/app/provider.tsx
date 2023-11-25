"use client";

import {createConnectTransport} from "@connectrpc/connect-web";
import {TransportProvider} from "@connectrpc/connect-query";
import {QueryClient, QueryClientProvider} from "@tanstack/react-query";

export default function Provider({ children }: { children: React.ReactNode}) {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        refetchInterval: 500,
      },
    }});

  const transport = createConnectTransport(
  {
    baseUrl: "http://localhost:3030/api/",
  });
  return (
    <QueryClientProvider client={queryClient}>
      <TransportProvider transport={transport}>
        {children}
      </TransportProvider>
    </QueryClientProvider>
  );
}
