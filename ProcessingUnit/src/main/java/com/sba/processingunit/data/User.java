package com.sba.processingunit.data;

import lombok.Data;
import org.springframework.data.annotation.Id;

import java.io.Serializable;

@Data
public class User implements Serializable {
    @Id
    private final String id;
    private final String name;
    private final double balance;
}
