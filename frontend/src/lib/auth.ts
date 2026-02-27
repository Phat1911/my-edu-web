const USER_KEY = '__eh_u'

export interface User {
  id: number
  username: string
  email: string
  display_name: string
}

function encode(value: string): string {
  if (typeof window === 'undefined') return value
  try { return btoa(encodeURIComponent(value)) } catch { return value }
}

function decode(value: string): string {
  if (typeof window === 'undefined') return value
  try { return decodeURIComponent(atob(value)) } catch { return value }
}

export function saveUser(user: User): void {
  if (typeof window === 'undefined') return
  const safe: User = {
    id: user.id,
    username: user.username,
    email: user.email,
    display_name: user.display_name,
  }
  sessionStorage.setItem(USER_KEY, encode(JSON.stringify(safe)))
}

export function getUser(): User | null {
  if (typeof window === 'undefined') return null
  const raw = sessionStorage.getItem(USER_KEY)
  if (!raw) return null
  try {
    return JSON.parse(decode(raw)) as User
  } catch {
    removeUser()
    return null
  }
}

export function removeUser(): void {
  if (typeof window === 'undefined') return
  sessionStorage.removeItem(USER_KEY)
}

export function isLoggedIn(): boolean {
  return getUser() !== null
}

export async function fetchMe(): Promise<User | null> {
  const apiUrl = process.env.NEXT_PUBLIC_API_URL ?? 'http://localhost:8080/api/v1'
  try {
    const res = await fetch(`${apiUrl}/auth/me`, {
      credentials: 'include',
    })
    if (!res.ok) {
      removeUser()
      return null
    }
    const data = await res.json()
    const user = data.user as User
    saveUser(user)
    return user
  } catch {
    return null
  }
}

// logout() calls the backend to clear the httpOnly cookie, then clears local user cache.
export async function logout(): Promise<void> {
  const apiUrl = process.env.NEXT_PUBLIC_API_URL ?? 'http://localhost:8080/api/v1'
  try {
    await fetch(`${apiUrl}/auth/logout`, {
      method: 'POST',
      credentials: 'include',
    })
  } catch {
    // best-effort - still clear local state even if request fails
  }
  removeUser()
}
