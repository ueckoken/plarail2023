import MapView from "./components/mapview/mapview"
import {TransportProvider} from "@connectrpc/connect-query";
import {createConnectTransport} from "@bufbuild/connect-web";
import {QueryClient, QueryClientProvider} from "@tanstack/react-query";
import React, {Suspense} from "react";
class ErrorBoundary extends React.Component {
  state: { hasError: boolean; };
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  constructor(props: { children: React.ReactNode; }) {
    super(props);
    this.state = { hasError: false };
  }

  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  static getDerivedStateFromError(error: any) {    // Update state so the next render will show the fallback UI.    
    console.log(error);
    return { hasError: true };  
  }
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  componentDidCatch(error: any, errorInfo: any) {    // You can also log the error to an error reporting service   
    console.log(error, errorInfo);
  }
  render() {
    if (this.state.hasError) {      // You can render any custom fallback UI      
      return <h1>Something went wrong.</h1>;
    }
    return this.props.children; 
  }
}

function App() {
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        refetchInterval: 500,
      },
    }});

  const transport = createConnectTransport(
  {
    baseUrl: "https://k1hmbp.tail5590e.ts.net/api/",
  });
  return (
  <>
    <ErrorBoundary>
      <Suspense>
        <QueryClientProvider client={queryClient}>
          <TransportProvider transport={transport}>
            <MapView/>
          </TransportProvider>
        </QueryClientProvider>
      </Suspense>
    </ErrorBoundary>
  </>
  );
}

export default App
