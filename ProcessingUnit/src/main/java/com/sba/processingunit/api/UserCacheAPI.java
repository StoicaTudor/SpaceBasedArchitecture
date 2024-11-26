package com.sba.processingunit.api;

import com.hazelcast.client.HazelcastClient;
import com.hazelcast.client.config.ClientConfig;
import com.hazelcast.core.HazelcastInstance;
import com.sba.processingunit.config.HazelcastUserConfig;
import org.apache.zookeeper.server.admin.CommandResponse;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.util.UUID;
import java.util.concurrent.ConcurrentMap;

@RestController
public class UserCacheAPI {
    private final HazelcastUserConfig hazelcastInstance;

    @Autowired
    public UserCacheAPI(HazelcastUserConfig hazelcastInstance) {
        this.hazelcastInstance = hazelcastInstance;
    }

    private ConcurrentMap<String, String> retrieveMap() {
        return hazelcastInstance.getInstance().getMap(HazelcastUserConfig.USERS);
    }

    @PostMapping("/put")
    public CommandResponse put(@RequestParam(value = "key") String key, @RequestParam(value = "value") String value) {
        retrieveMap().put(key, value);

        return new CommandResponse(value);
    }

    @GetMapping("/x")
    public void get() {
//        hazelcastInstance.getInstance()
//                .getMap(HazelcastUserConfig.USERS)
//                .put(UUID.randomUUID().toString(), UUID.randomUUID().toString());

        ClientConfig clientConfig = new ClientConfig();
        new ClientConfig().getNetworkConfig().addAddress("127.0.0.1");

        HazelcastInstance x = HazelcastClient.newHazelcastClient(clientConfig);
        x.getMap(HazelcastUserConfig.USERS).put(UUID.randomUUID().toString(), UUID.randomUUID().toString());
    }

    @GetMapping("/get")
    public CommandResponse get(@RequestParam(value = "key") String key) {
        String value = retrieveMap().get(key);
        return new CommandResponse(value);
    }
}