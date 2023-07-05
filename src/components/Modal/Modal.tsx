'use client'
import { ReactNode, useEffect, useRef } from "react"

type ModalProps = {
    children: ReactNode
    tWidth: string
    tHeight: string
    isOpen: boolean
    setIsOpen: React.Dispatch<React.SetStateAction<boolean>>
    goBack: () => void
    isSubPage: boolean
}

const Modal = ({ children, tWidth, tHeight, isOpen, setIsOpen, goBack, isSubPage }: ModalProps) => {
    const modalRef = useRef<HTMLDivElement | null>(null)

    const closeModal = async () => {
        if (isOpen && modalRef.current) {
            modalRef.current.classList.add("opacity-0")
            await sleep(150)
            setIsOpen(false)
            goBack()
        }
    }

    const openModal = async () => {
        if (isOpen && modalRef.current) {
            await sleep(1)
            modalRef.current.classList.remove("opacity-0")
        }
    }

    useEffect(() => {
        if (isOpen && modalRef.current) {
            openModal()
        }
    }, [isOpen])


    return (
        <>
            {isOpen &&

                <div className="absolute left-0 top-0 w-full h-full z-20 flex justify-center items-center transition opacity-0" ref={modalRef}>
                    {/* Transparent background */}
                    <div className="absolute w-full h-full opacity-50 bg-black" onClick={() => closeModal()}></div>

                    {/* Content */}
                    <div className={`bg-foreground z-30 rounded-lg ${tWidth} ${tHeight} flex justify-center items-center p-2 transition-all`}>

                        {/* Container */}
                        <div className="flex flex-col w-full h-full">

                            {/* Top bar. Go back and close modal */}
                            <div className="flex w-full justify-between">
                                {isSubPage ?
                                    <span onClick={() => goBack()} className="cursor-pointer">
                                        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" strokeWidth={2} stroke="currentColor" className="w-6 h-6">
                                            <path strokeLinecap="round" strokeLinejoin="round" d="M19.5 12h-15m0 0l6.75 6.75M4.5 12l6.75-6.75" />
                                        </svg>
                                    </span>
                                    :
                                    <div></div>
                                }

                                <span onClick={() => closeModal()} className="cursor-pointer">
                                    <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" strokeWidth={2} className="w-6 h-6 stroke-red-400 transition hover:stroke-red-600">
                                        <path strokeLinecap="round" strokeLinejoin="round" d="M6 18L18 6M6 6l12 12" />
                                    </svg>
                                </span>
                            </div>

                            {children}
                        </div>
                    </div>
                </div>
            }
        </>
    )
}

export default Modal

const sleep = (ms: number) => new Promise(res => setTimeout(res, ms));