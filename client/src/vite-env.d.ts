/// <reference types="vite/client" />
interface ImportMetaEnv {
    readonly SERVER_BASE_URL: string
}

interface ImportMeta {
    readonly env: ImportMetaEnv
}
