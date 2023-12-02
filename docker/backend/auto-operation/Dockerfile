FROM node:21.3-bookworm AS base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

FROM base AS dep-resolver

WORKDIR /auto-operation
COPY --chown=node:node pnpm-lock.yaml pnpm-workspace.yaml ./
COPY backend/auto-operation ./backend/auto-operation
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile

FROM dep-resolver AS auto-operation-builder
RUN pnpm run --filter=./backend/auto-operation -r build
RUN pnpm deploy --filter=./backend/auto-operation --prod dist/auto-operation
RUN cp backend/auto-operation/main.js dist/auto-operation/main.js
WORKDIR /autooperation/dist/auto-operation


FROM gcr.io/distroless/nodejs20-debian12:nonroot AS dashboard-runner
COPY --from=auto-operation-builder --chown=nonroot:nonroot /auto-operation/dist/auto-operation /auto-operation
ENV SERVER_ADDR=''
CMD ["/auto-operation/main.js"]

