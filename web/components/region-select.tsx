import { Dispatch, SetStateAction } from 'react'
import {
    Select,
    SelectContent,
    SelectItem,
    SelectTrigger,
    SelectValue,
} from './ui/select'
import { Region } from '@/app/dashboard/monitors/data-table'
import ReactCountryFlag from 'react-country-flag'

interface RegionSelectProps {
    regions: Region[]
    region: Region
    setRegion: Dispatch<SetStateAction<Region>>
}

export function RegionSelect({
    regions,
    region,
    setRegion,
}: RegionSelectProps) {
    return (
        <Select
            value={region.regionName}
            onValueChange={(value) =>
                setRegion(regions.find((r) => r.regionName == value)!)
            }
        >
            <SelectTrigger
                className="flex w-40 **:data-[slot=select-value]:block **:data-[slot=select-value]:truncate"
                size="sm"
                aria-label="Select a value"
            >
                <SelectValue placeholder="Select Region" />
            </SelectTrigger>
            <SelectContent>
                {regions.map((r) => (
                    <SelectItem
                        key={r.regionId}
                        value={r.regionName}
                        className="flex items-center"
                    >
                        <div className="flex items-center gap-2">
                            <ReactCountryFlag
                                countryCode={r.regionName.toUpperCase()}
                                svg
                            />
                            {r.regionName}
                        </div>
                    </SelectItem>
                ))}
            </SelectContent>
        </Select>
    )
}
