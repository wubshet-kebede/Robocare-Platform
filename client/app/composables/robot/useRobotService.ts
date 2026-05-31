export const useRobotService = () => {
  const { $api } = useNuxtApp();

  const getRobots = () => {
    return $api("/get-robots", {
      method: "GET",
    });
  };

  return { getRobots };
};
