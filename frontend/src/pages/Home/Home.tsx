import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { healthCheck, listFiles, queryFile } from "../../api";
import FileList from "../../components/FileList/FileList";
import FileUploader from "../../components/FileUploader/FileUploader";
import Search from "../../components/Search/Search";
import "./Home.css";

function Home() {
  const [FileHash, setFileHash] = useState<string | null>(null);
  const [healthCheckStatus, setHealthCheckStatus] = useState<string | null>(
    null
  );
  const [snippets, setSnippets] = useState<string[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [userFiles, setUserFiles] = useState<any>([]);

  const navigate = useNavigate()

  useEffect(() => {
    const checkHealth = async () => {
      try {
        await healthCheck();
        setHealthCheckStatus("Backend is healthy");
      } catch (error) {
        setHealthCheckStatus("Backend is not healthy");
        console.error("Health check failed:", error);
        setError(
          "Failed to connect to the backend. Please ensure it is running."
        );
      }
    };
    const listUserFiles = async () => {
      try {
        const files = await listFiles();
        setUserFiles(files);
      } catch (error) {
        setError("Failed to get user files");
      }
    };
    checkHealth();
    listUserFiles();
  }, []);

  const handleUpload = (name: string) => {
    listFiles().then((files) => {
      setUserFiles(files);
    });
  };

  const handleSearch = async (query: string) => {
    if (!FileHash) {
      setError('Please upload or select a file to search in')
      return;
    }

    try {
      const results = await queryFile(FileHash, query);
      setSnippets(results);
      if (results.length === 0) setError("No results found.");
    } catch (err) {
      console.error("Search error:", err);
    }
  };

  return (
    <div className="App">
      <header>
        <div className="header-text">
          <h1>Welcome to the App</h1>
          <p>
            This is a simple application that lets you upload a text file and
            search it.
          </p>
          {healthCheckStatus && (
            <p
              className="health-status"
              style={{
                color: healthCheckStatus.includes("not") ? "red" : "green",
              }}
            >
              {healthCheckStatus}
            </p>
          )}
          <p>Selected File: {FileHash || "None"}</p>
        </div>
        <FileUploader onUpload={handleUpload} setError={setError} />
        <button className="logout-button" onClick={()=>{
          localStorage.removeItem('token');
          navigate('/login')
        }}> Logout </button>
      </header>

      <div className="error-message">{error && <p>{error}</p>}</div>

      <div className="main">
        <div className="search-container">
          <Search onSearch={handleSearch} snippets={snippets} />
        </div>
        <FileList
          files={userFiles ?? []}
          selectedFile={FileHash}
          onSelect={setFileHash}
        />
      </div>
    </div>
  );
}

export default Home;
