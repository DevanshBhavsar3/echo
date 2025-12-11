'use client'

import { motion } from 'motion/react'
import { useState } from 'react'

type Metrics = { name: string; value: string }

const metrics: Metrics[] = [
    { name: 'Status', value: 'Up' },
    { name: 'Response Time', value: '297 MS' },
    { name: 'Status Code', value: '200 OK' },
]

const LINE_COLOR = 'stroke-neutral-300'
const GRADIENT_COLOR = 'stroke-cyan-500'

const lineVariant = {
    hidden: {
        pathLength: 0,
    },
    visible: {
        pathLength: 1,
        transition: {
            duration: 1,
            delay: 1,
        },
    },
}

function LineSVG() {
    return (
        <motion.svg
            initial={'hidden'}
            animate={'visible'}
            exit={'hidden'}
            width="100%"
            viewBox="0 0 1440 794"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
        >
            <defs>
                <linearGradient id="gradient">
                    <stop
                        offset={0.55}
                        stopColor="white"
                        stopOpacity={0}
                    ></stop>
                    <stop
                        offset={0.65}
                        stopColor="white"
                        stopOpacity={1}
                    ></stop>
                    <stop
                        offset={0.85}
                        stopColor="white"
                        stopOpacity={1}
                    ></stop>
                    <stop offset={1} stopColor="white" stopOpacity={0}></stop>
                </linearGradient>

                <mask id="gradient-mask" maskUnits={'userSpaceOnUse'}>
                    <motion.rect
                        initial={{
                            x: '150%',
                        }}
                        animate={{
                            x: '-100%',
                        }}
                        transition={{
                            delay: 2.5,
                            duration: 2,
                            repeat: Infinity,
                            ease: 'linear',
                        }}
                        y={0}
                        width={'120%'}
                        height={'100%'}
                        fill="url(#gradient)"
                    />
                </mask>
            </defs>

            {/*Bottom Line*/}
            <motion.path
                variants={lineVariant}
                d="M1439.94 428.49C809.627 435.027 463.104 490.945 47.9999 698"
                className={`${LINE_COLOR} stroke-2`}
            />
            <path
                d="M1439.94 428.49C809.627 435.027 463.104 490.945 47.9999 698"
                className={`${GRADIENT_COLOR} stroke-2`}
                mask="url(#gradient-mask)"
            />

            {/*Top Line*/}
            <motion.path
                variants={lineVariant}
                d="M1440.05 374.859C807.688 368.984 460.123 311.212 44 96"
                className={`${LINE_COLOR} stroke-2`}
            />
            <path
                d="M1440.05 374.859C807.688 368.984 460.123 311.212 44 96"
                className={`${GRADIENT_COLOR} stroke-2`}
                mask="url(#gradient-mask)"
            />

            {/*Middle Line*/}
            <motion.line
                variants={lineVariant}
                x1="1440"
                y1="401"
                x2="96"
                y2="401"
                className={`${LINE_COLOR} stroke-2`}
            />
            <line
                x1="96"
                y1="401"
                x2="1440"
                y2="401"
                className={`${GRADIENT_COLOR} stroke-2`}
                mask="url(#gradient-mask)"
            />

            {/*Icons*/}
            <motion.rect
                variants={lineVariant}
                x="1"
                y="1"
                width="94"
                height="94"
                rx="17"
                className="fill-neutral-50 stroke-zinc-200 stroke-2"
            />
            <motion.rect
                variants={lineVariant}
                x="1"
                y="355"
                width="94"
                height="94"
                rx="17"
                className="fill-neutral-50 stroke-zinc-200 stroke-2"
            />
            <motion.rect
                variants={lineVariant}
                x="1"
                y="699"
                width="94"
                height="94"
                rx="17"
                className="fill-neutral-50 stroke-zinc-200 stroke-2"
            />
            <motion.svg
                initial={{
                    y: 50,
                    opacity: 0,
                }}
                animate={{
                    y: 0,
                    opacity: 1,
                }}
                transition={{
                    delay: 0.5,
                }}
            >
                <path
                    opacity="0.987"
                    fillRule="evenodd"
                    clipRule="evenodd"
                    d="M44.9337 22.2257C47.7238 21.9248 49.3748 21.9248 52.3727 22.2257C57.6795 22.9841 62.5989 25.3526 66.4196 28.9887C63.8377 31.3451 61.2898 33.7359 58.7766 36.1604C53.9634 32.2215 48.591 31.3124 42.6593 33.4329C38.308 35.3652 35.2779 38.4964 33.5692 42.8267C30.777 40.8195 28.021 38.7655 25.3028 36.6657C25.1139 36.5697 24.8981 36.5346 24.687 36.5654C29.005 28.5267 35.7526 23.7455 44.9298 22.222"
                    fill="#F44336"
                />
                <path
                    opacity="0.997"
                    fillRule="evenodd"
                    clipRule="evenodd"
                    d="M24.6795 36.5653C24.8976 36.5331 25.1041 36.5666 25.2991 36.6656C28.0174 38.7654 30.7733 40.8194 33.5656 42.8266C33.1262 44.5139 32.8492 46.2367 32.7381 47.9731C32.8331 49.6527 33.1089 51.3014 33.5656 52.919L24.8873 59.5891C21.1081 51.964 21.0389 44.2894 24.6795 36.5653Z"
                    fill="#FFC107"
                />
                <path
                    opacity="0.999"
                    fillRule="evenodd"
                    clipRule="evenodd"
                    d="M66.0072 67.6674C63.3051 65.3664 60.4762 63.208 57.533 61.2017C60.4834 59.1901 62.2743 56.4305 62.9054 52.9226H48.4468V43.2279C56.7851 43.161 65.1195 43.2291 73.4501 43.4322C75.0306 51.7187 73.2051 59.1901 67.9738 65.8466C67.3518 66.485 66.6928 67.0927 66.0072 67.6674Z"
                    fill="#448AFF"
                />
                <path
                    opacity="0.993"
                    fillRule="evenodd"
                    clipRule="evenodd"
                    d="M33.5669 52.9227C36.7227 60.4957 42.5082 64.0308 50.9234 63.5279C53.2857 63.2639 55.5506 62.467 57.5351 61.2018C60.4804 63.2133 63.3052 65.3685 66.0093 67.6674C61.7247 71.385 56.2617 73.5909 50.5116 73.925C49.2053 74.0259 47.8926 74.0259 46.5862 73.925C36.7906 72.8103 29.5581 68.0316 24.8887 59.5891L33.5669 52.9227Z"
                    fill="#43A047"
                />
                <path d="M48 381L74 424H22L48 381Z" fill="black" />
                <mask
                    id="mask0_0_1"
                    maskUnits="userSpaceOnUse"
                    x="22"
                    y="720"
                    width="52"
                    height="52"
                >
                    <path d="M22 720H74V772H22V720Z" fill="white" />
                </mask>
                <g mask="url(#mask0_0_1)">
                    <path
                        fillRule="evenodd"
                        clipRule="evenodd"
                        d="M48 723.939C33.635 723.939 22 735.574 22 749.939C22 761.444 29.4425 771.162 39.7775 774.607C41.0775 774.834 41.565 774.054 41.565 773.372C41.565 772.754 41.5325 770.707 41.5325 768.529C35 769.732 33.31 766.937 32.79 765.474C32.4975 764.727 31.23 762.419 30.125 761.802C29.215 761.314 27.915 760.112 30.0925 760.079C32.14 760.047 33.6025 761.964 34.09 762.744C36.43 766.677 40.1675 765.572 41.6625 764.889C41.89 763.199 42.5725 762.062 43.32 761.412C37.535 760.762 31.49 758.519 31.49 748.574C31.49 745.747 32.4975 743.407 34.155 741.587C33.895 740.937 32.985 738.272 34.415 734.697C34.415 734.697 36.5925 734.014 41.565 737.362C43.645 736.777 45.855 736.484 48.065 736.484C50.275 736.484 52.485 736.777 54.565 737.362C59.5375 733.982 61.715 734.697 61.715 734.697C63.145 738.272 62.235 740.937 61.975 741.587C63.6325 743.407 64.64 745.714 64.64 748.574C64.64 758.552 58.5625 760.762 52.7775 761.412C53.72 762.224 54.5325 763.784 54.5325 766.222C54.5325 769.699 54.5 772.494 54.5 773.372C54.5 774.054 54.9875 774.867 56.2875 774.607C61.4489 772.864 65.9339 769.547 69.1113 765.122C72.2887 760.697 73.9984 755.387 74 749.939C74 735.574 62.365 723.939 48 723.939Z"
                        fill="black"
                    />
                </g>
            </motion.svg>
        </motion.svg>
    )
}

export function Background() {
    const [hovered, setHovered] = useState(false)

    return (
        <div className="absolute inset-0 select-none">
            <motion.div
                initial={{
                    y: 50,
                    opacity: 0,
                }}
                animate={{
                    y: 0,
                    opacity: 1,
                }}
                transition={{
                    delay: 0.5,
                }}
                onHoverStart={() => setHovered(true)}
                onHoverEnd={() => setHovered(false)}
                className="group absolute bottom-10 left-1/2 z-10 flex w-52 -translate-x-1/2 flex-col items-center gap-5 rounded-sm border bg-white/40 p-4 backdrop-blur-xl"
            >
                <div className="w-full text-left">
                    <span className="flex items-center gap-3 font-mono text-xs text-neutral-500">
                        URL <p className="font-mono text-black">google.com</p>
                    </span>
                    <div className="w-full border-spacing-1.5 border border-dashed" />
                </div>
                <span className="w-full text-left font-mono text-xs text-neutral-500">
                    METRICS
                    <div className="w-full border-spacing-1.5 border border-dashed" />
                </span>
                <div className="flex w-full flex-col gap-3">
                    {metrics.map((metric, i) => (
                        <div
                            key={metric.name}
                            className={`flex w-full items-center ${hovered ? 'justify-between' : 'justify-center'} gap-3`}
                        >
                            <motion.span
                                layoutId={`metric-title-${metric.name}`}
                                className="font-mono text-sm font-medium"
                                transition={{
                                    type: 'spring',
                                    layout: {
                                        delay: i * 0.15,
                                    },
                                }}
                            >
                                {metric.name}
                            </motion.span>
                            {hovered && (
                                <motion.p
                                    initial={{
                                        filter: 'blur(10px)',
                                    }}
                                    animate={{
                                        filter: 'blur(0px)',
                                    }}
                                    className="text-muted-foreground font-mono text-xs"
                                    transition={{
                                        delay: i * 0.3,
                                    }}
                                >
                                    {metric.value}
                                </motion.p>
                            )}
                        </div>
                    ))}
                </div>
            </motion.div>

            <div className="absolute bottom-0 w-1/2">
                <LineSVG />
            </div>

            <div className="absolute bottom-0 left-1/2 w-1/2 -scale-y-100 rotate-180">
                <LineSVG />
            </div>
        </div>
    )
}
