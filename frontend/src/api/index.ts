const BASE_URL = "http://localhost:8080";

// health check endpoint
export async function healthCheck(): Promise<void> {
  const res = await fetch(`${BASE_URL}/ping`);
  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.error || "Health check failed");
  }
}

/**
 * Uploads a text file to the backend.
 * @param file File object to upload.
 */
export async function uploadFile(file: File): Promise<void> {
  const formData = new FormData();
  formData.append("file", file);

  const res = await fetch(`${BASE_URL}/upload`, {
    method: "POST",
    headers: {
      "Authorization": `${localStorage.getItem("token")}`,
    },
    body: formData,
  });

  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.error || "Failed to upload file");
  }
}

/**
 * Uploads multiple text files to the backend.
 * @param files Array of File objects to upload.
 */
export async function uploadFiles(files: File[]): Promise<void> {
  const formData = new FormData();

  files.forEach((file) => {
    formData.append("files", file);
  });

  const res = await fetch(`${BASE_URL}/uploads`, {
    method: "POST",
    headers: {
      Authorization: `${localStorage.getItem("token")}`,
    },
    body: formData,
  });

  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.error || "Failed to upload files");
  }
}


/**
 * Sends a search query to the backend for a specific file.
 * @param FileHash FileHash of the uploaded file.
 * @param query Search term.
 * @returns Array of matching text snippets.
 */
export async function queryFile(FileHash: string, query: string): Promise<string[]> {
  const res = await fetch(`${BASE_URL}/query`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      "Authorization": `${localStorage.getItem("token")}`,
    },
    body: JSON.stringify({ FileHash, query }),
  });

  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.error || "Failed to query file");
  }

  const data = await res.json();
  return data.results;
}

export async function login(username: string, password: string): Promise<void> {
  const res = await fetch(`${BASE_URL}/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ username, password }),
  });

  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.error || "Login failed");
  }

  const data = await res.json();
  localStorage.setItem("token", data.token);
}

export async function signup(username: string, password: string): Promise<void> {
  const res = await fetch(`${BASE_URL}/signup`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ username, password }),
  });

  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.error || "Signup failed");
  }

  const data = await res.json();
  localStorage.setItem("token", data.token);
}

export async function listFiles(): Promise<string[]> {
  const res = await fetch(`${BASE_URL}/files`, {
    method: "GET",
    headers: {
      "Authorization": `${localStorage.getItem("token")}`,
    },
  });

  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.error || "Failed to list files");
  }

  const data = await res.json();
  console.log("Files:", data);
  return data;
}

export interface IQueryHistory {
  "ID": string,
  "Query": string,
  "FileHash": string,
  "UserId": string,
  "Timestamp": number
}

export async function getQueryHistory(): Promise<IQueryHistory[]>{
  const res = await fetch(`${BASE_URL}/query/history`, {
    method: "GET",
    headers: {
      "Authorization": `${localStorage.getItem("token")}`,
    },
  });

  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.error || "Failed to get query history");
  }

  const data = await res.json();
  return data;
}

export interface IUserProfile {
    "created_at": number,
    "username": string
}

export async function getUserProfile(): Promise<IUserProfile>{
  const res = await fetch(`${BASE_URL}/profile`, {
    method: "GET",
    headers: {
      "Authorization": `${localStorage.getItem("token")}`,
    },
  });

  if (!res.ok) {
    const err = await res.json();
    throw new Error(err.error || "Failed to get user profile");
  }

  const data = await res.json();
  return data;
}
