export const useRobotService = () => {
  const { $api } = useNuxtApp();

  const getRobots = () => {
    return $api("/get-robots", {
      method: "GET",
    });
  };
  const publishNavGoal = (data: {
    patient_id: string;
    robot_id: string;
    room_id: string;
  }) => {
    return $api("/publish-nav-goal", {
      method: "POST",
      body: data,
    });
  };

  return { getRobots, publishNavGoal };
};
