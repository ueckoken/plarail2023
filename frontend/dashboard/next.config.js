/** @type {import('next').NextConfig} */
const nextConfig = {
    output: "standalone",
    sassOptions: { prependData: `@import "@/app/_variables";` },

    webpack: (config) => {
        config.module.rules.push({
            test: /\.svg$/,
            use: [
                {
                    loader: "@svgr/webpack",
                },
            ],
        });
        return config;
    },
    images: {
        disableStaticImages: true, // importした画像の型定義設定を無効にする
    },
};

module.exports = nextConfig;
