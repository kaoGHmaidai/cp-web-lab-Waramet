import { NextRequest, NextResponse } from "next/server";

export async function GET(request: NextRequest) {
  const name = request.nextUrl.searchParams.get("name");
  const query = name ? `?name=${encodeURIComponent(name)}` : "";

  const response = await fetch(`${process.env.BACKEND_URL}/greet${query}`);

  if (!response.ok) {
    return NextResponse.json(
      { message: "Request failed" },
      { status: response.status },
    );
  }

  const data = await response.json();
  return NextResponse.json(data);
}
