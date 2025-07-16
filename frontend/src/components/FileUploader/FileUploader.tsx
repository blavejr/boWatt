import React, { useState } from "react";
import { uploadFiles } from "../../api";
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
    const files = Array.from(e.dataTransfer.files).filter(isTxtFile);

    if (files.length === 0) {
      setError("Only .txt files are allowed.");
      return;
    }

    setError(null);
    try {
      await uploadFiles(files);
      files.forEach(file => onUpload(file.name));
    } catch (err: any) {
      setError(err.message);
    }
  };

  const handleFileInput = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const files = Array.from(e.target.files || []).filter(isTxtFile);

    if (files.length === 0) {
      setError("Only .txt files are allowed.");
      return;
    }

    setError(null);
    try {
      await uploadFiles(files);
      files.forEach(file => onUpload(file.name));
    } catch (err: any) {
      setError(err.message);
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
      <p>Drag and drop .txt files here, or</p>
      <input type="file" accept=".txt" multiple onChange={handleFileInput} />
    </div>
  );
}
