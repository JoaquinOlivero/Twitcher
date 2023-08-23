export function debounce<Params extends any[]>(
    func: (...args: Params) => any,
    timeout: number = 300,
): (...args: Params) => void {
    let timer: number
    return (...args: Params) => {
        clearTimeout(timer)
        timer = window.setTimeout(() => {
            func(...args)
        }, timeout)
    }
}