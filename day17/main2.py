from collections import defaultdict

def modify_x_velocity(velocity):
    if velocity > 0:
        return velocity - 1
    elif velocity < 0:
        return velocity + 1
    else:
        return velocity

def modify_y_velocity(velocity):
    return velocity - 1

def is_x_hopeless(start_velocity, current_velocity, current_position, target_range, positions, step_max):
    if (step_max > 0 and len(positions) >= step_max) or (current_position > max(target_range)):
        # print(f"Starting with x velocity of {start_velocity} ran out of x after {len(positions)} steps: {positions}")
        return True

def is_y_hopeless(start_velocity, current_velocity, current_position, target_range, positions, step_max):
    if (step_max > 0 and len(positions) >= step_max) or (current_position < min(target_range)):
        # print(f"Starting with y velocity of {start_velocity} overshot y after {len(positions)} steps: {positions}")
        return True

def will_reach(axis_name, start_position, start_velocity, target_range, step_modifier, is_hopeless, step_max=0):
    # print(f"Trying start {axis_name} velocity of {start_velocity}")
    current_position = start_position
    positions = []
    current_velocity = start_velocity
    valid_velocities = {}
    while True:
        if current_position in target_range:
            # print(f"Starting with {axis_name} velocity of {start_velocity} reached {current_position} in {len(positions)} steps: {positions}")
            valid_velocities[len(positions)] = max(positions)

        if is_hopeless(start_velocity, current_velocity, current_position, target_range, positions, step_max):
            # print(f"Starting with {axis_name} velocity of {start_velocity} failed at {current_position} in {len(positions)} steps: {positions}")
            return valid_velocities

        current_position += current_velocity
        positions.append(current_position)
        current_velocity = step_modifier(current_velocity)


if __name__ == '__main__':
    # start = 0,0
    start_x_position = 0
    start_y_position = 0
    # real: target area: x=128..160, y=-142..-88
    target_x_range = range(128, 160+1)
    target_y_range = range(-142, -88+1)

    # test: target area: x=20..30, y=-10..-5
    # target_x_range = range(20, 30+1)
    # target_y_range = range(-10, -5+1)

    y_steps = defaultdict(list)
    for y_velocity in range(min(target_y_range), 141 + 1):
        valid_y_velocities = will_reach('y', start_y_position, y_velocity, target_y_range, modify_y_velocity, is_y_hopeless)
        for steps,max_value in valid_y_velocities.items():
            y_steps[steps].append(y_velocity)

    x_steps = defaultdict(list)
    for x_velocity in range(0, max(target_x_range) + 1):
        valid_x_velocities = will_reach('x', start_x_position, x_velocity, target_x_range, modify_x_velocity, is_x_hopeless, max(y_steps))
        for steps, max_value in valid_x_velocities.items():
            x_steps[steps].append(x_velocity)

    common_steps = [step for step in y_steps if step in x_steps]

    start_velocities = []
    for step in common_steps:
        ys = y_steps[step]
        xs = x_steps[step]
        for x in xs:
            for y in ys:
                start_velocities.append((x, y))

    print(start_velocities)
    print(len(set(start_velocities)))
