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
RUN pnpm deploy --filter=./frontend/dashboard --prod --legacy dist/dashboard

FROM nginx:alpine AS dashboard-runner
COPY --from=dashboard-builder /dashboard/frontend/dashboard/dist /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
