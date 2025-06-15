import MapView from "./components/mapview/mapview"
import { TransportProvider } from "@connectrpc/connect-query";
import { createConnectTransport } from "@bufbuild/connect-web";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import React, { Suspense } from "react";
class ErrorBoundary extends React.Component {
  state: { hasError: boolean; };
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  constructor(props: { children: React.ReactNode; }) {
    super(props);
    this.state = { hasError: false };
  }

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  render() {
    if (this.state.hasError) {      // You can render any custom fallback UI      
      return <h1>Something went wrong.</h1>;
    }
    return (this.props as { children: React.ReactNode }).children;
  }
}

function App() {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        refetchInterval: 500,
      },
    }
  });

  const transport = createConnectTransport(
    {
      baseUrl: "https://api-plarail2023.ueckoken.club/api/",
    });
  return (
    <>
      <ErrorBoundary>
        <Suspense>
          <QueryClientProvider client={queryClient}>
            <TransportProvider transport={transport}>
              <MapView />
            </TransportProvider>
          </QueryClientProvider>
        </Suspense>
      </ErrorBoundary>
    </>
  );
}

export default App
