export const useAuthUser = () => {
  return useState<any>("user", () => null);
};
