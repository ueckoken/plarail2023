import React, { Suspense } from "react";

export const metadata = { title: "ホーム" };

const Layout: React.FC<React.PropsWithChildren> = ({ children }) => (
    <Suspense>{children}</Suspense>
);

export default Layout;
