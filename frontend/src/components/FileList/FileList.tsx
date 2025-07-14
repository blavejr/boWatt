import React from "react";
import './FileList.css';

interface FileListProps {
  files: [{
    ID: string,
    FileHash: string,
    UserId: string,
    Name: string,
    Timestamp: string,
    Content: string
  }];
  selectedFile: string | null;
  onSelect: (FileHash: string) => void;
}

const FileList: React.FC<FileListProps> = ({ files, selectedFile, onSelect }) => {
  return (
    <div className="file-list">
      <h3>Uploaded Files</h3>
      <ul>
        {files.map((file, index) => (
          <li
            key={index}
            className={file?.FileHash === selectedFile ? "selected" : ""}
            onClick={() => onSelect(file?.FileHash)}
          >
            <p>File Name: {file?.Name}</p>
            <p>{file.Content.substring(0, 100)}...</p>
            <p>Uploaded At: {new Date(file?.Timestamp).toLocaleString()}</p>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default FileList;
