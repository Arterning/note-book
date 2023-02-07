import Cookies from 'js-cookie'

const TokenKey = 'Blockchain-Platform-Token'
const NameKey = 'Blockchain-Admin-Name'
// const UsernameKey = 'Blockchain-Admin-Username'
const UserIdKey = 'Blockchain-Platform-UserId'

const pwdStatusKey = 'Blockchain-Admin-PwdStatus'

export function getToken() {
  return Cookies.get(TokenKey)
}

export function setToken(token) {
  return Cookies.set(TokenKey, token)
}

export function setName(name) {
  return Cookies.set(NameKey, name)
}

export function removeToken() {
  return Cookies.remove(TokenKey)
}

export function getUserId() {
  return Cookies.get(UserIdKey)
}

export function setUserId(userId) {
  return Cookies.set(UserIdKey, userId)
}

export function setPwdStatus(status) {
  return Cookies.set(pwdStatusKey, status)
}
