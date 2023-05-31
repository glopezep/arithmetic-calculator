
INSERT INTO public.users (id, created_at, updated_at, deleted_at, email, "password", balance) VALUES
('808a8952-47de-40f7-8107-ae798e5d6e2e', '2023-05-30 21:47:03.568158+00', '2023-05-30 21:47:03.568158+00', NULL, 'test@gmail.com', '$2a$10$4DHm/gS6hFgK45zju1IgVe99RUa9zxj55OIa54WRA9v/d8vXpvPR2', 100);


INSERT INTO public.operations (id, created_at, updated_at, deleted_at, "type", cost) VALUES
    ('34050426-838d-4f28-b181-530a76400286', '2023-05-30 21:47:03.568158+00', '2023-05-30 21:47:03.568158+00', NULL, 'addition', 10),
    ('ae7d8f38-05bf-44a5-ad4d-967a7d5669ad', '2023-05-30 21:47:03.568158+00', '2023-05-30 21:47:03.568158+00', NULL, 'subtraction', 10),
    ('70feb8f3-125e-47d0-96d3-d680ca26c33b', '2023-05-30 21:47:03.568158+00', '2023-05-30 21:47:03.568158+00', NULL, 'multiplication', 10),
    ('c3d23fac-4b2f-42d9-ae77-6c94056c5d40', '2023-05-30 21:47:03.568158+00', '2023-05-30 21:47:03.568158+00', NULL, 'division', 10),
    ('6774e63d-6011-48ea-80d3-856d9f07cfc9', '2023-05-30 21:47:03.568158+00', '2023-05-30 21:47:03.568158+00', NULL, 'square_root', 10),
    ('c081fbe8-bcf2-4376-949f-a5848c5cfb4e', '2023-05-30 21:47:03.568158+00', '2023-05-30 21:47:03.568158+00', NULL, 'random_string', 20);