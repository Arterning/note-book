import Cookies from 'js-cookie'

const TokenKey = 'Blockchain-Admin-Token'
const NameKey = 'Blockchain-Admin-Name'
const UsernameKey = 'Blockchain-Admin-Username'
const pwdStatusKey = 'Blockchain-Admin-PwdStatus'

export function getToken(domain, tokenKey) {
  return domain
    ? Cookies.get(tokenKey || TokenKey, { domain })
    : Cookies.get(tokenKey || TokenKey)
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

export function getUsernameLTS() {
  return localStorage.getItem(UsernameKey)
}

export function setUsernameLTS(name) {
  return localStorage.setItem(UsernameKey, name)
}

export function getPwdStatus() {
  return Cookies.get(pwdStatusKey)
}

export function setPwdStatus(status) {
  return Cookies.set(pwdStatusKey, status)
}

export function removePwdStatus() {
  return Cookies.remove(pwdStatusKey)
}
