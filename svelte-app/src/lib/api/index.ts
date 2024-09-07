export * from './users'
export * from './api'

export function cleanUrlParams(params: Record<string, any>): string {
  const cleanParams = Object.fromEntries(
    Object.entries(params).filter(([_, value]) => value !== undefined)
  );
  return new URLSearchParams(cleanParams as Record<string, string>).toString();
}
