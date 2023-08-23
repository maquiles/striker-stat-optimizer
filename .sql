SELECT * FROM striker 
WHERE shooting > strength
    AND shooting > speed
    AND shooting > passing
    AND shooting > technique
    AND speed > strength
    AND speed > passing
    AND speed > technique
ORDER BY shooting DESC, speed DESC;