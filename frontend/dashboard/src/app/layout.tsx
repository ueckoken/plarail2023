import type { Metadata } from "next";
import { Inter } from "next/font/google";

import "./globals.scss";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
    title: {
        template: "%s - 工学研究部 プラレール展示 管理画面",
        default: "工学研究部 プラレール展示 管理画面",
    },
    description: "電気通信大学 調布祭 プラレール企画",
};

export default function RootLayout({
    children,
}: {
    children: React.ReactNode;
}) {
    return (
        <html lang="ja">
            <body className={inter.className}>{children}</body>
        </html>
    );
}
