interface MobileToggleProps {
  isOpen: boolean;
  onToggle: () => void;
}

export default function MobileToggle({ isOpen, onToggle }: MobileToggleProps) {
  if (isOpen) return null;

  return (
    <button
      onClick={onToggle}
      className="lg:hidden fixed top-4 left-4 z-50 p-2 bg-blue-500 text-white rounded-lg shadow-lg cursor-pointer transition-all duration-200 hover:scale-110 hover:bg-blue-600"
    >
      <span className="material-symbols-outlined" style={{ fontVariationSettings: "'wght' 700" }}>
        menu
      </span>
    </button>
  );
}
