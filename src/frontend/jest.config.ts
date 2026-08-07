/** @type {import('ts-jest').JestConfigWithTsJest} */
const jestConfig = {
  preset: "ts-jest",
  testEnvironment: "jsdom",
  clearMocks: true,
  setupFilesAfterEnv: ["<rootDir>/jest.setup.ts"],
  coverageReporters: ["json-summary", "text", "lcov"],
  moduleNameMapper: {
    "^@/(.*)$": "<rootDir>/$1",
  },
};

export default jestConfig;
