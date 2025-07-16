import { fireEvent, render, screen } from "@testing-library/react";
import Search from "./Search";

describe("Search Component", () => {
  it("renders input and responds to typing", () => {
    const mockOnSearch = jest.fn();
    render(<Search value="testing" onSearch={mockOnSearch} snippets={[]} />);

    const input = screen.getByPlaceholderText("Search...");

    fireEvent.change(input, { target: { value: "hello" } });

    expect(input).toHaveValue("hello");
    expect(mockOnSearch).toHaveBeenCalledWith("hello");
  });

  it("displays 'No results' when snippets is empty", () => {
    render(<Search value="testing" onSearch={() => {}} snippets={[]} />);
    expect(screen.getByText("No results")).toBeInTheDocument();
  });

  it("renders provided snippets", () => {
    const snippets = ["Line 1", "Line 2"];
    render(<Search value="testing" onSearch={() => {}} snippets={snippets} />);

    snippets.forEach((snippet) => {
      expect(screen.getByText(snippet)).toBeInTheDocument();
    });
  });
});
