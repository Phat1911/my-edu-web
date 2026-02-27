import api from "./api"

const USER_KEY = '__eh_u'

export interface User {
  id: number
  username: string
  email: string
  display_name: string
}

// Token lives in an httpOnly cookie set by the backend - JS never touches it.
// Only non-sensitive display data (id, username, email, display_name) is cached
// in localStorage so the UI can read it without a round-trip on every page load.

export function saveUser(user: User): void {
  if (typeof window === 'undefined') return
  const safe: User = {
    id: user.id,
    username: user.username,
    email: user.email,
    display_name: user.display_name,
  }
  localStorage.setItem(USER_KEY, JSON.stringify(safe))
}

export function getUser(): User | null {
  if (typeof window === 'undefined') return null
  const raw = localStorage.getItem(USER_KEY)
  if (!raw) return null
  try {
    return JSON.parse(raw) as User
  } catch {
    removeUser()
    return null
  }
}

export function removeUser(): void {
  if (typeof window === 'undefined') return
  localStorage.removeItem(USER_KEY)
}

export function isLoggedIn(): boolean {
  return getUser() !== null
}

// fetchMe() is the source of truth for auth state.
// It hits /auth/me which reads the httpOnly cookie automatically.
// On success it refreshes the localStorage cache and returns the user.
// On failure it clears the stale cache and returns null.
export async function fetchMe(): Promise<User | null> {
  try {
    const data = await api.get <{user: User}>(`/auth/me`)
    const user = await data.user
    saveUser(user)
    return user
  } catch {
    return null
  }
}

// logout() clears the httpOnly cookie via the backend, then removes the localStorage cache.
export async function logout(): Promise<void> {
  try {
    await api.post(`/auth/logout`);
  } catch {
    // best-effort - still clear local state even if the request fails
  }
  removeUser()
}
