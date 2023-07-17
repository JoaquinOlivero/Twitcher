type Props = {
    text: string
    waitingText?: string
    width: string
    disabled: boolean | undefined
    isWaiting: boolean
    backgroundColor: string
    backgroundColorHover: string
    onClick: () => Promise<void>
}

const ActionButton = ({ text, waitingText, width, onClick, disabled, isWaiting, backgroundColor, backgroundColorHover }: Props) => {
    return (
        <div onClick={() => onClick()} className={`flex justify-center items-center ${backgroundColor} w-${width} py-1 rounded font-semibold tracking-wider capitalize cursor-pointer transition hover:${backgroundColorHover} ${disabled && 'pointer-events-none opacity-40 '} ${isWaiting && 'pointer-events-none '}`}>
            {isWaiting ?
                <div className="flex justify-center items-center gap-2 h-6">
                    <span>{waitingText}</span>
                    <svg className="animate-spin h-5 w-5 text-white" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                        <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4"></circle>
                        <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"></path>
                    </svg>
                </div>
                :
                <button className={`font-semibold tracking-wider capitalize cursor-pointer transition ${disabled && 'pointer-events-none opacity-40'}`}>{text}</button>
            }
        </div>
    )
}

export default ActionButton