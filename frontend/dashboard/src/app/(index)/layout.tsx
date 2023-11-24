import React, { Suspense } from "react";
import Provider from "@/app/provider";

export const metadata = { title: "ホーム" };

const Layout: React.FC<React.PropsWithChildren> = ({ children }) => (
    <Provider>
      <Suspense>
        {children}
      </Suspense>
    </Provider>
);

export default Layout;
