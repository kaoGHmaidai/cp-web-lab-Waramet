import { fireEvent, render, screen, waitFor } from "@testing-library/react";
import { GreetForm } from "@/components/greet-form";

describe("GreetForm", () => {
  beforeEach(() => {
    global.fetch = jest.fn();
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  it("renders the form", () => {
    render(<GreetForm />);

    expect(
      screen.getByRole("heading", { name: "Say hello" }),
    ).toBeInTheDocument();
    expect(screen.getByLabelText("Your name")).toBeInTheDocument();
    expect(
      screen.getByRole("button", { name: "Get greeting" }),
    ).toBeInTheDocument();
  });

  it("submits the name and displays the returned message", async () => {
    (global.fetch as jest.Mock).mockResolvedValueOnce({
      ok: true,
      json: async () => ({ message: "Hello, Ann!" }),
    });

    render(<GreetForm />);

    fireEvent.change(screen.getByLabelText("Your name"), {
      target: { value: "Ann" },
    });
    fireEvent.click(screen.getByRole("button", { name: "Get greeting" }));

    await waitFor(() => {
      expect(screen.getByRole("status")).toHaveTextContent("Hello, Ann!");
    });

    expect(global.fetch).toHaveBeenCalledWith("/api/greet?name=Ann");
  });

  it("shows an error state when the request fails", async () => {
    (global.fetch as jest.Mock).mockResolvedValueOnce({ ok: false });

    render(<GreetForm />);

    fireEvent.click(screen.getByRole("button", { name: "Get greeting" }));

    await waitFor(() => {
      expect(screen.getByRole("alert")).toHaveTextContent(
        "Something went wrong. Try again.",
      );
    });
  });
});
