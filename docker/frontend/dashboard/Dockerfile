FROM node:20.10-bookworm AS base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

FROM base AS dep-resolver

WORKDIR /dashboard
COPY --chown=node:node pnpm-lock.yaml pnpm-workspace.yaml ./
COPY frontend/ ./frontend 
RUN pnpm install --frozen-lockfile

FROM dep-resolver AS dashboard-builder
WORKDIR /dashboard
RUN pnpm run --filter=./frontend/dashboard build
RUN pnpm deploy --filter=./frontend/dashboard --prod dist/dashboard
WORKDIR /dashboard/dist/dashboard
RUN cp -r .next/static .next/standalone/.next/
RUN cp -r public .next/standalone/

FROM gcr.io/distroless/nodejs20-debian12:nonroot AS dashboard-runner
COPY --from=dashboard-builder --chown=nonroot:nonroot /dashboard/dist/dashboard /dashboard
CMD ["/dashboard/.next/standalone/server.js"]
