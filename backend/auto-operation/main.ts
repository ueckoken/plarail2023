import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@bufbuild/connect-web";
import { StateManagerService } from "./proto/state/v1/state_connectweb";


async function main() {
  const transport = createConnectTransport(
    {
      // baseUrl: process.env.NEXT_PUBLIC_API_ENDPOINT!,
      baseUrl: "http://localhost:8080",
    });
  const client = createPromiseClient(StateManagerService, transport);

  const res = await client.getPointStates({})
  console.log(res)
}

main()