export function HeaderLine({ children }: { children?: React.ReactNode }) {
  return (
    <div className="w-full bg-primary/50 py-2 px-2 text-balance text-center text-xs font-medium">
      {children}
    </div>
  );
}
