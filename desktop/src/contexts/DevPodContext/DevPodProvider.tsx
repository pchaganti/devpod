import { useQuery } from "@tanstack/react-query"
import { createContext, ReactNode, useMemo } from "react"
import { client } from "../../client"
import { QueryKeys } from "../../queryKeys"
import { TProInstances, TProviders, TQueryResult } from "../../types"
import { REFETCH_INTERVAL_MS, REFETCH_PROVIDER_INTERVAL_MS } from "./constants"
import { usePollWorkspaces } from "./workspaces"
import { useChangeSettings } from "../SettingsContext"

export type TDevpodContext = Readonly<{
  providers: TQueryResult<TProviders>
  proInstances: TQueryResult<TProInstances>
}>
export const DevPodContext = createContext<TDevpodContext>(null!)

export function DevPodProvider({ children }: Readonly<{ children?: ReactNode }>) {
  const { set } = useChangeSettings()
  usePollWorkspaces()

  const providersQuery = useQuery({
    queryKey: QueryKeys.PROVIDERS,
    queryFn: async () => (await client.providers.listAll()).unwrap(),
    refetchInterval: REFETCH_PROVIDER_INTERVAL_MS,
  })

  const proInstancesQuery = useQuery({
    queryKey: QueryKeys.PRO_INSTANCES,
    queryFn: async () => {
      const proInstances = (await client.pro.listAll({ authenticate: true })).unwrap()
      if (proInstances !== undefined && proInstances.length > 0) {
        set("experimental_devPodPro", true)
      }

      return proInstances
    },
    refetchInterval: REFETCH_INTERVAL_MS,
  })

  const value = useMemo<TDevpodContext>(
    () => ({
      providers: [
        providersQuery.data,
        { status: providersQuery.status, error: providersQuery.error },
      ],
      proInstances: [
        proInstancesQuery.data,
        { status: proInstancesQuery.status, error: proInstancesQuery.error },
      ],
    }),
    [
      providersQuery.data,
      providersQuery.status,
      providersQuery.error,
      proInstancesQuery.data,
      proInstancesQuery.status,
      proInstancesQuery.error,
    ]
  )

  return <DevPodContext.Provider value={value}>{children}</DevPodContext.Provider>
}
