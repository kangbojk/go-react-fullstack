export function eraseCookie(name) {
    document.cookie = name + '=; Max-Age=0'
}