import { render, screen, waitFor } from "@testing-library/react";
import App from "./App";
import * as api from "./api";

// Mock the API functions
jest.mock("./api");

describe("App Component", () => {
  beforeEach(() => {
    jest.clearAllMocks();
  });

  it("renders the welcome message and health status", async () => {
    // Mock healthCheck to succeed
    (api.healthCheck as jest.Mock).mockResolvedValueOnce(undefined);

    render(<App />);

    expect(screen.getByText(/Welcome to the App/i)).toBeInTheDocument();
    expect(screen.getByText(/upload a text file/i)).toBeInTheDocument();

    // Wait for health check to complete
    await waitFor(() => {
      expect(screen.getByText(/Backend is healthy/i)).toBeInTheDocument();
    });
  });
});
