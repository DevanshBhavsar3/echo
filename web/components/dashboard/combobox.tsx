'use client'

import { Check, ChevronDown } from 'lucide-react'
import { Popover, PopoverContent, PopoverTrigger } from '../ui/popover'
import { Button } from '../ui/button'
import { useState } from 'react'
import { Command, CommandGroup, CommandItem, CommandList } from '../ui/command'

export interface ComboboxItem {
    value: string
    label: string
}

export function Combobox({ data }: { data: ComboboxItem[] }) {
    const [open, setOpen] = useState(false)
    const [value, setValue] = useState(data[0]?.value || '')

    return (
        <Popover open={open} onOpenChange={setOpen}>
            <PopoverTrigger asChild>
                <Button
                    variant="outline"
                    role="combobox"
                    aria-expanded={open}
                    className="w-50 justify-between"
                >
                    {value
                        ? data.find((item) => item.value === value)?.label
                        : 'Select an option'}
                    <ChevronDown className="opacity-50" />
                </Button>
            </PopoverTrigger>
            <PopoverContent className="w-50 p-0">
                <Command>
                    <CommandList>
                        <CommandGroup>
                            {data.map((item: ComboboxItem) => (
                                <CommandItem
                                    key={item.value}
                                    value={item.value}
                                    onSelect={(value) => {
                                        setValue(value)
                                        setOpen(false)
                                    }}
                                >
                                    {item.label}
                                    <Check
                                        className={`ml-auto ${item.value === value ? 'opacity-100' : 'opacity-0'}`}
                                    />
                                </CommandItem>
                            ))}
                        </CommandGroup>
                    </CommandList>
                </Command>
            </PopoverContent>
        </Popover>
    )
}
