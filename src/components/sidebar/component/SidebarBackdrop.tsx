interface SidebarBackdropProps {
  isOpen: boolean;
  onClose: () => void;
}

export default function SidebarBackdrop({ isOpen, onClose }: SidebarBackdropProps) {
  if (!isOpen) return null;

  return (
    <div
      onClick={onClose}
      className="lg:hidden fixed inset-0 bg-black/50 z-40"
    />
  );
}
