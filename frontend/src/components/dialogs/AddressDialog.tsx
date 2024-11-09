import { useState } from "react";

export function InputDialog({ isOpen, onClose }: { isOpen: boolean; onClose: () => void }) {
  const [inputValue, setInputValue] = useState("");

  const handleSubmit = () => {
    try {
      // Handle the result as needed
      onClose();
    } catch (error) {
      console.error("Error:", error);
    }
  };

  if (!isOpen) return null;

  return (
    <div className="modal">
      <input type="text" value={inputValue} onChange={(e) => setInputValue(e.target.value)} />
      <button onClick={handleSubmit}>Submit</button>
      <button onClick={onClose}>Cancel</button>
    </div>
  );
}

export default InputDialog;
