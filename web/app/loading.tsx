import {
    Empty,
    EmptyDescription,
    EmptyHeader,
    EmptyMedia,
    EmptyTitle,
} from '@/components/ui/empty'
import { Spinner } from '@/components/ui/spinner'

export default function Loading() {
    return (
        <div className="flex h-screen w-full items-center justify-center">
            <Empty className="w-full">
                <EmptyHeader>
                    <EmptyMedia variant="icon">
                        <Spinner />
                    </EmptyMedia>
                    <EmptyTitle>Loading</EmptyTitle>
                    <EmptyDescription>
                        Please wait while we process your request. Do not
                        refresh the page.
                    </EmptyDescription>
                </EmptyHeader>
            </Empty>
        </div>
    )
}
