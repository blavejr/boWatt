import React, { useState } from "react";
import { uploadFile } from "../../api";
import './FileUploader.css';

interface FileUploaderProps {
  onUpload: (filename: string) => void;
  setError: (msg: string | null) => void;
}

export default function FileUploader({ onUpload, setError }: FileUploaderProps) {
  const [dragOver, setDragOver] = useState(false);
  const isTxtFile = (file: File) => file.name.toLowerCase().endsWith(".txt");

  const handleDrop = async (e: React.DragEvent) => {
    e.preventDefault();
    setDragOver(false);
    const file = e.dataTransfer.files[0];
    if (file) {
      if (!isTxtFile(file)) {
        setError("Only .txt files are allowed.");
        return;
      }
      setError(null);
      await uploadFile(file);
      onUpload(file.name);
    }
  };

  const handleFileInput = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0];
    if (file) {
      if (!isTxtFile(file)) {
        setError("Only .txt files are allowed.");
        return;
      }
      setError(null);
      await uploadFile(file);
      onUpload(file.name);
    }
  };

  return (
    <div
      className={`upload-box ${dragOver ? "drag-over" : ""}`}
      onDrop={handleDrop}
      onDragOver={(e) => {
        e.preventDefault();
        setDragOver(true);
      }}
      onDragLeave={() => setDragOver(false)}
    >
      <p>Drag and drop a .txt file here, or</p>
      <input type="file" onChange={handleFileInput} />
    </div>
  );
}
