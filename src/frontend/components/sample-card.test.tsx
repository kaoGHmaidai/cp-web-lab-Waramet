import { render, screen } from "@testing-library/react";

import { SampleCard } from "./sample-card";

describe("SampleCard", () => {
  it("renders the supplied content", () => {
    render(
      <SampleCard
        title="Ship with confidence"
        description="This example proves the Jest setup can render React components."
        actionLabel="Run checks"
      />,
    );

    expect(
      screen.getByRole("heading", { name: "Ship with confidence" }),
    ).toBeInTheDocument();
    expect(
      screen.getByText(
        "This example proves the Jest setup can render React components.",
      ),
    ).toBeInTheDocument();
    expect(
      screen.getByRole("button", { name: "Run checks" }),
    ).toBeInTheDocument();
  });

  it("uses default actionLabel when not provided", () => {
    render(<SampleCard title="T" description="D" />);
    expect(
      screen.getByRole("button", { name: "Learn more" }),
    ).toBeInTheDocument();
  });
});
